import { NgModule } from '@angular/core';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { ReactiveFormsModule } from '@angular/forms';
import { FormlyModule } from '@ngx-formly/core';
import { FormlyBootstrapModule } from '@ngx-formly/bootstrap';
import { NgxSpinnerModule } from 'ngx-spinner';

import { SignupComponent } from './components/signup/signup.component';
import { SigninComponent } from './components/signin/signin.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ToastrModule } from 'ngx-toastr';
import { TokenInterceptor } from './intercepters/token.interceptor';
import { VerifyAccountComponent } from './components/verify-account/verify-account.component';
import { ProfileComponent } from './components/profile/profile.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { WarehouseComponent } from './components/warehouse/warehouse.component';
import { ContainerComponent } from './components/container/container.component';
import { GeneralComponent } from './components/profile/general/general.component';
import { DocumentsComponent } from './components/profile/documents/documents.component';
import { WhGeneralComponent } from './components/warehouse/wh-general/wh-general.component';
import { WhInfraComponent } from './components/warehouse/wh-infra/wh-infra.component';
import { WhSurroundingComponent } from './components/warehouse/wh-surrounding/wh-surrounding.component';

@NgModule({
   declarations: [
      AppComponent,
      SignupComponent,
      SigninComponent,
      PageNotFoundComponent,
      DashboardComponent,
      VerifyAccountComponent,
      ProfileComponent,
      SidebarComponent,
      WarehouseComponent,
      ContainerComponent,
      GeneralComponent,
      DocumentsComponent,
      WhGeneralComponent,
      WhInfraComponent,
      WhSurroundingComponent,
   ],
   imports: [
      BrowserModule,
      AppRoutingModule,
      HttpClientModule,
      NgbModule,
      ReactiveFormsModule,
      FormlyModule.forRoot(),
      FormlyBootstrapModule,
      BrowserAnimationsModule, // required animations module
      ToastrModule.forRoot(), // ToastrModule added
      NgxSpinnerModule,
   ],
   providers: [
      { provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi: true },
   ],
   bootstrap: [AppComponent],
})
export class AppModule {}
