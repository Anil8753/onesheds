import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ContainerComponent } from './components/container/container.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { ProfileComponent } from './components/profile/profile.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { SigninComponent } from './components/signin/signin.component';
import { SignupComponent } from './components/signup/signup.component';
import { VerifyAccountComponent } from './components/verify-account/verify-account.component';
import { WarehouseComponent } from './components/warehouse/warehouse.component';

const routes: Routes = [
   { path: 'signin', component: SigninComponent },
   { path: 'signup', component: SignupComponent },
   { path: 'verify_account', component: VerifyAccountComponent },
   {
      path: 'container',
      component: ContainerComponent,
      children: [
         // {
         //    path: 'sidebar',
         //    component: SidebarComponent,
         // },
         {
            path: 'dashboard',
            component: DashboardComponent,
         },
         {
            path: 'profile',
            component: ProfileComponent,
         },
         {
            path: 'warehouse',
            component: WarehouseComponent,
         },
      ],
   },
   {
      path: '',
      redirectTo: '/dashboardng generate component page-not-found',
      pathMatch: 'full',
   },
   { path: '**', component: PageNotFoundComponent },
];

@NgModule({
   imports: [RouterModule.forRoot(routes)],
   exports: [RouterModule],
})
export class AppRoutingModule {}
