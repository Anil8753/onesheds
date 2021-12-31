import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { FormlyFieldConfig } from '@ngx-formly/core';
import { ConfigService } from 'src/app/services/config.service';
import { ToastrService } from 'ngx-toastr';
import { HttpClient } from '@angular/common/http';
import { AuthService } from 'src/app/services/auth.service';
import { ErrorMsgService } from 'src/app/services/error-msg.service';
import { INTERCEPTOR_NO_AUTH_HEADER } from '../../services/constants.service';

@Component({
   selector: 'app-signin',
   templateUrl: './signin.component.html',
   styleUrls: ['./signin.component.scss'],
})
export class SigninComponent implements OnInit {
   form = new FormGroup({});
   model: SigninModel = { email: '', password: '' };
   fields: FormlyFieldConfig[] = [];

   constructor(
      private http: HttpClient,
      private router: Router,
      private toastr: ToastrService,
      private configService: ConfigService,
      private authService: AuthService,
      private errMsgService: ErrorMsgService
   ) {
      this.initFormly();
   }

   ngOnInit(): void {}

   submit() {
      if (!this.form.valid) {
         return;
      }

      const postData = {
         user: this.model.email,
         password: this.model.password,
      };

      this.http
         .post<SigninResp>(
            `${this.configService.baseUrl()}/api/v1/signin`,
            postData,
            {
               headers: { [INTERCEPTOR_NO_AUTH_HEADER]: 'true' },
            }
         )
         .subscribe({
            next: (v) => {
               this.authService.setToken(v.Desc.access_token);
               this.router.navigateByUrl('dashboard');
            },
            error: (e) => {
               this.toastr.error(this.errMsgService.get(e.error), 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }

   private initFormly() {
      this.fields = [
         {
            key: 'email',
            type: 'input',
            templateOptions: {
               label: 'Email address',
               placeholder: 'Enter email',
               required: true,
            },
         },
         {
            key: 'password',
            type: 'input',
            templateOptions: {
               label: 'Password',
               placeholder: 'Enter password',
               required: true,
            },
         },
      ];
   }
}

interface SigninModel {
   email: string;
   password: string;
}

interface SigninResp {
   Code: number;
   CodeDesc: string;
   Desc: Desc;
}

interface Desc {
   access_token: string;
}
