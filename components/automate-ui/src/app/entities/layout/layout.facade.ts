import { Injectable } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';

import { LayoutSidebarService } from './layout-sidebar.service';
import { notificationState } from 'app/entities/notifications/notification.selectors';
import { Notification, Type } from 'app/entities/notifications/notification.model';

import * as fromLayout from './layout.reducer';
import { MenuItemGroup } from './layout.model';
import { sidebar, showPageLoading } from './layout.selectors';
import { ShowPageLoading } from './layout.actions';

// Important! These must match components/automate-ui/src/styles/_variables.scss
enum Height {
  Navigation = '70px',
  Banner = '110px'
}

export enum Sidebar {
  Dashboards = 'dashboards',
  Applications = 'applications',
  Infrastructure = 'infrastructure',
  Compliance = 'compliance',
  Settings = 'settings',
  Profile = 'profile'
}

@Injectable({
  providedIn: 'root'
})
export class LayoutFacadeService {
  headerHeight: string = Height.Navigation;
  contentHeight = `calc(100% - ${this.headerHeight})`;
  sidebar$: Observable<MenuItemGroup[]>;
  showPageLoading$: Observable<boolean>;
  showLicenseNotification = false;
  showHeader = true;
  enableSidebar = true;

  constructor(
    private store: Store<fromLayout.LayoutEntityState>,
    private layoutSidebarService: LayoutSidebarService
  ) {
    this.sidebar$ = store.select(sidebar);
    this.showPageLoading$ = store.select(showPageLoading);

    store.select(notificationState).subscribe(
      (notifications: Notification[]) => {
        this.showLicenseNotification =
          notifications &&  notifications.some(n => n.type === Type.license);
        this.updateContentHeight(
          this.showLicenseNotification ? Height.Banner : Height.Navigation);
    });
  }

  hasGlobalNotifications(): boolean {
    return this.headerHeight === Height.Banner;
  }

  ShowPageLoading(showLoading: boolean) {
    this.store.dispatch( new ShowPageLoading(showLoading));
  }

  showFullPage() {
    this.contentHeight = '100%';
    this.enableSidebar = false;
    this.showHeader = false;
  }

  hideFullPage(): void {
    this.updateContentHeight(this.headerHeight);
    this.enableSidebar = true;
    this.showHeader = true;
  }

  private updateContentHeight(height: string): void {
    this.headerHeight = height;
    this.contentHeight = `calc(100% - ${this.headerHeight})`;
  }

  showSidebar(sidebarName: string) {
    this.layoutSidebarService.updateSidebars(sidebarName);
  }
}
