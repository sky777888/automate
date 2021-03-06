import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MockComponent } from 'ng2-mock-component';
import { Observable, BehaviorSubject } from 'rxjs';
import { StoreModule, Store } from '@ngrx/store';
import { NgrxStateAtom, ngrxReducers, defaultInitialState, runtimeChecks } from 'app/ngrx.reducers';

import { LicenseNotificationsComponent } from './license-notifications.component';
import { LicenseFacadeService } from 'app/entities/license/license.facade';
import { Notification } from 'app/entities/notifications/notification.model';

describe('LicenseNotificationsComponent', () => {
  let store: Store<NgrxStateAtom>;
  let component: LicenseNotificationsComponent;
  let fixture: ComponentFixture<LicenseNotificationsComponent>;

  class MockLicenseFacadeService extends LicenseFacadeService {
    notifications$: Observable<Notification[]> = new BehaviorSubject([]);
  }

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        StoreModule.forRoot(ngrxReducers, { initialState: defaultInitialState, runtimeChecks })
      ],
      declarations: [
        MockComponent({ selector: 'chef-notification', inputs: ['type', 'timeout'] }),
        LicenseNotificationsComponent],
      providers: [{ provide: LicenseFacadeService, useClass:  MockLicenseFacadeService }]
    }).compileComponents();
      store = TestBed.get(Store);
      spyOn(store, 'dispatch').and.callThrough();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LicenseNotificationsComponent);
    component = fixture.componentInstance;
    store = TestBed.get(Store);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
