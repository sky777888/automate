import { Component, EventEmitter, OnInit, OnDestroy } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Store, select } from '@ngrx/store';
import { interval as observableInterval,  Observable, Subject, combineLatest } from 'rxjs';
import { map, takeUntil, filter, take } from 'rxjs/operators';
import { isNil } from 'lodash/fp';

import { NgrxStateAtom } from 'app/ngrx.reducers';
import { Regex } from 'app/helpers/auth/regex';
import { ChefSorters } from 'app/helpers/auth/sorter';
import { HttpStatus } from 'app/types/types';
import { loading, EntityStatus } from 'app/entities/entities';
import { LayoutFacadeService } from 'app/entities/layout/layout.facade';
import { ProjectStatus } from 'app/entities/rules/rule.model';
import { isIAMv2 } from 'app/entities/policies/policy.selectors';
import { ProjectService } from 'app/entities/projects/project.service';
import {
  allProjects, getAllStatus, createStatus, createError
} from 'app/entities/projects/project.selectors';
import { GetProjects, CreateProject, DeleteProject  } from 'app/entities/projects/project.actions';
import { Project } from 'app/entities/projects/project.model';
import { ApplyRulesStatus, ApplyRulesStatusState } from 'app/entities/projects/project.reducer';
import { LoadOptions } from 'app/services/projects-filter/projects-filter.actions';
import { ChefKeyboardEvent } from 'app/types/material-types';

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.scss']
})
export class ProjectListComponent implements OnInit, OnDestroy {
  public isIAMv2$: Observable<boolean>;
  public sortedProjects$: Observable<Project[]>;
  public projectToDelete: Project;
  public deleteModalVisible = false;
  public messageModalVisible = false;
  public createModalVisible = false;
  public createProjectForm: FormGroup;
  public creatingProject = false;
  public conflictErrorEvent = new EventEmitter<boolean>();
  public confirmApplyStartModalVisible = false;
  public confirmApplyStopModalVisible = false;

  // This flag governs filling the above cache.
  // The state returned by this.projects.applyRulesStatus$ (Running, NotRunning)
  // is not available soon enough--we need to know the instant the user starts the update.
  private applyRulesInProgress = false;

  // True if there are any rules that have a status of 'EDITS_PENDING'.
  private projectsHaveStagedChanges = false;

  private percentageComplete = 0;

  private updateProjectsFailed = false;
  private updateProjectsCancelled = false;
  public cancelRulesInProgress = false;

  public applyRulesButtonText$: Observable<string>;
  public ApplyRulesStatusState = ApplyRulesStatusState;

  private isDestroyed = new Subject<boolean>();

  public statusLabel: Record<ProjectStatus, string> = {
      'PROJECT_RULES_STATUS_UNSET': '',
      'NO_RULES': 'No rules',
      'EDITS_PENDING': 'Edits pending',
      'RULES_APPLIED': 'Applied'
  };

  constructor(
    private layoutFacade: LayoutFacadeService,
    private store: Store<NgrxStateAtom>,
    public projects: ProjectService,
    fb: FormBuilder
  ) {
    this.createProjectForm = fb.group({
      // Must stay in sync with error checks in create-object-modal.component.html
      name: ['', Validators.required],
      id: ['',
        [Validators.required, Validators.pattern(Regex.patterns.ID), Validators.maxLength(64)]]
    });
  }

  ngOnInit(): void {
    this.layoutFacade.showSettingsSidebar();
    this.projects.getApplyRulesStatus();
    this.store.dispatch(new GetProjects());
    this.store.pipe(
      select(getAllStatus),
      takeUntil(this.isDestroyed),
      map(loading)
    ).subscribe((isLoading) =>
      this.layoutFacade.ShowPageLoading(isLoading)
    );
    this.sortedProjects$ = this.store.select(allProjects).pipe(
      map((unsorted: Project[]) => ChefSorters.naturalSort(unsorted, 'name')));

    this.isIAMv2$ = this.store.select(isIAMv2);

    this.projects.applyRulesStatus$
      .pipe(takeUntil(this.isDestroyed))
      .subscribe(({ state, failed, cancelled, percentageComplete }: ApplyRulesStatus) => {
        if (this.applyRulesInProgress && state === ApplyRulesStatusState.NotRunning) {
          this.cancelRulesInProgress = false;
          this.percentageComplete = 0;
          this.closeConfirmApplyStopModal();
        }
        this.applyRulesInProgress = state === ApplyRulesStatusState.Running;
        this.updateProjectsFailed = failed;
        this.updateProjectsCancelled = cancelled;
        if (!this.cancelRulesInProgress && state === ApplyRulesStatusState.Running) {
          this.percentageComplete = percentageComplete;
        }
      });

    this.store.select(allProjects).pipe(
      takeUntil(this.isDestroyed)
    ).subscribe((projectList: Project[]) => {
      this.projectsHaveStagedChanges = projectList.some(p => p.status === 'EDITS_PENDING');
    });

    // handle project creation success response
    this.store.pipe(
      select(createStatus),
      takeUntil(this.isDestroyed),
      filter(state => {
        return this.creatingProject && state === EntityStatus.loadingSuccess;
      }))
      .subscribe(() => {
        this.creatingProject = false;
        this.closeCreateModal();

        // This is issued periodically from projects-filter.effects.ts; we do it now
        // so the user doesn't have to wait.
        this.store.dispatch(new LoadOptions());
      });

    combineLatest([
      this.store.select(createStatus),
      this.store.select(createError)
    ]).pipe(
      takeUntil(this.isDestroyed),
      filter(() => this.createModalVisible),
      filter(([state, error]) => state === EntityStatus.loadingFailure && !isNil(error)))
      .subscribe(([_, error]) => {
        if (error.status === HttpStatus.CONFLICT) {
          this.conflictErrorEvent.emit(true);
          this.creatingProject = false;
        } else {
          // close modal on any error other than conflict and display in banner
          this.closeCreateModal();
        }
      });
  }

  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }

  public closeDeleteModal(): void {
    this.deleteModalVisible = false;
  }

  public closeMessageModal(): void {
    this.messageModalVisible = false;
  }

  public startProjectDelete($event: ChefKeyboardEvent, p: Project): void {
    if ($event.isUserInput) {
      const deletableStates: ProjectStatus[] = ['EDITS_PENDING', 'RULES_APPLIED'];
      if (deletableStates.includes(p.status)) {
        this.messageModalVisible = true;
      } else {
        this.projectToDelete = p;
        this.deleteModalVisible = true;
      }
    }
  }

  public deleteProject(): void {
    this.closeDeleteModal();
    this.store.dispatch(new DeleteProject({id: this.projectToDelete.id}));
  }

  public createProject(): void {
    this.creatingProject = true;
    const project = {
      id: this.createProjectForm.controls['id'].value,
      name: this.createProjectForm.controls['name'].value.trim()
    };
    this.store.dispatch(new CreateProject(project));
  }

  public openCreateModal(): void {
    this.createModalVisible = true;
    this.resetCreateModal();
  }

  public closeCreateModal(): void {
    this.createModalVisible = false;
    this.resetCreateModal();
  }

  resetCreateModal(): void {
    this.creatingProject = false;
    this.createProjectForm.reset();
    this.conflictErrorEvent.emit(false);
  }

  public openConfirmUpdateStartModal(): void {
    this.confirmApplyStartModalVisible = true;
  }

  private closeConfirmApplyStartModal(): void {
    this.confirmApplyStartModalVisible = false;
  }

  public confirmApplyStart(): void {
    this.closeConfirmApplyStartModal();
    this.projects.applyRulesStart();
    this.applyRulesInProgress = true;

    // Rapid sampling for 3 seconds for more responsive UX.
    // If the update is still running, the secondary (active) emitter
    // will check this status at frequent intervals.
    // Once the update completes, the tertiary (dormant) emitter
    // will check this status at INfrequent intervals.
    // (See getActiveApplyRulesStatus$ and getDormantApplyRulesStatus$.)
    observableInterval(250).pipe(take(12)) // 12 x 250ms => 3 seconds
      .subscribe(() => {
        this.projects.getApplyRulesStatus();
      });
  }

  public cancelApplyStart(): void {
    this.closeConfirmApplyStartModal();
  }

  public openConfirmUpdateStopModal(): void {
    this.confirmApplyStopModalVisible = true;
  }

  private closeConfirmApplyStopModal(): void {
    this.confirmApplyStopModalVisible = false;
  }

  public confirmApplyStop(): void {
    this.cancelRulesInProgress = true;
    this.projects.applyRulesStop();
  }

  public cancelApplyStop(): void {
    this.closeConfirmApplyStopModal();
  }

  public getButtonText(): string {
    if (this.applyRulesInProgress) {
      return `Updating Projects ${Math.round(this.percentageComplete * 100)}%...`;
    }
    if (this.projectsHaveStagedChanges
      || this.updateProjectsCancelled
      || this.updateProjectsFailed) {
      return 'Update Projects';
    }
    return 'Projects Up-to-Date';
  }

  public isDisabled(): boolean {
    return this.applyRulesInProgress ||
      (!this.projectsHaveStagedChanges
        && !this.updateProjectsCancelled
        && !this.updateProjectsFailed);
  }

}
