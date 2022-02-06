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

import { SignupComponent } from './components/signup/signup.component';
import { SigninComponent } from './components/signin/signin.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ToastrModule } from 'ngx-toastr';
import { TokenInterceptor } from './intercepters/token.interceptor';
import { VerifyAccountComponent } from './components/verify-account/verify-account.component';
import { ListingComponent } from './components/listing/listing.component';
import { BookingComponent } from './components/booking/booking.component';
import { OrdersComponent } from './components/orders/orders.component';

@NgModule({
   declarations: [
      AppComponent,
      SignupComponent,
      SigninComponent,
      PageNotFoundComponent,
      DashboardComponent,
      VerifyAccountComponent,
      ListingComponent,
      BookingComponent,
      OrdersComponent,
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
   ],
   providers: [
      { provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi: true },
   ],
   bootstrap: [AppComponent],
})
export class AppModule {}
