import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { FormlyFieldConfig } from '@ngx-formly/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';
import { UserRegistrationData, DummyRegData } from '../types';

@Component({
   selector: 'app-general',
   templateUrl: './general.component.html',
   styleUrls: ['./general.component.scss'],
})
export class GeneralComponent implements OnInit {
   @Input()
   data!: UserRegistrationData;

   model = DummyRegData();
   form = new FormGroup({});
   fields: FormlyFieldConfig[] = [];

   constructor(
      private configService: ConfigService,
      private http: HttpClient,
      private spinner: NgxSpinnerService,
      private toastr: ToastrService
   ) {}

   ngOnInit(): void {
      this.model = this.data;
      this.initFormly();
   }

   submit() {
      if (!this.form.valid) {
         return;
      }

      this.spinner.show();

      this.http
         .put<any>(`${this.configService.baseUrl()}/api/v1/profile`, this.model)
         .subscribe({
            next: (v) => {
               console.log(v);
            },
            error: (e) => {
               this.toastr.error(
                  'Failed to update the profile data.',
                  'Error!'
               );
               console.error(e);
            },
            complete: () => this.spinner.hide(),
         });
   }

   private initFormly() {
      this.fields = [
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-6',
                  type: 'input',
                  key: 'firstName',
                  templateOptions: {
                     label: 'First Name',
                     required: true,
                  },
               },
               {
                  className: 'col-6',
                  type: 'input',
                  key: 'lastName',
                  templateOptions: {
                     label: 'Last Name',
                     required: true,
                  },
                  expressionProperties: {
                     'templateOptions.disabled': '!model.firstName',
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-6',
                  key: 'email',
                  type: 'input',
                  templateOptions: {
                     label: 'Email',
                     placeholder: 'Enter email',
                     type: 'email',
                     required: true,
                     readonly: true,
                  },
               },
               {
                  className: 'col-6',
                  key: 'phone',
                  type: 'input',
                  templateOptions: {
                     label: 'Phone number',
                     placeholder: 'Enter phone number',
                     type: 'phone',
                     required: true,
                  },
               },
            ],
         },
         {
            key: 'address',
            type: 'textarea',
            templateOptions: {
               label: 'Address',
               placeholder: 'Enter Address',
               rows: 2,
               required: true,
            },
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-3',
                  key: 'pincode',
                  type: 'input',
                  templateOptions: {
                     label: 'PIN',
                     placeholder: 'Enter PIN number',
                     required: true,
                  },
               },
               {
                  className: 'col-3',
                  key: 'city',
                  type: 'input',
                  templateOptions: {
                     label: 'City',
                     placeholder: 'Enter city',
                     required: true,
                  },
               },
               {
                  className: 'col-3',
                  key: 'district',
                  type: 'input',
                  templateOptions: {
                     label: 'District',
                     placeholder: 'Enter district',
                     required: true,
                  },
               },
               {
                  className: 'col-3',
                  key: 'state',
                  type: 'input',
                  templateOptions: {
                     label: 'State',
                     placeholder: 'Enter state',
                     required: true,
                  },
               },
            ],
         },
      ];
   }
}
