import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { FormlyFieldConfig } from '@ngx-formly/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';
import { DummyRegData, UserRegistrationData } from '../types';

@Component({
   selector: 'app-documents',
   templateUrl: './documents.component.html',
   styleUrls: ['./documents.component.scss'],
})
export class DocumentsComponent implements OnInit {
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
                  key: 'pancard',
                  type: 'input',
                  templateOptions: {
                     label: 'PAN card',
                     placeholder: 'Enter PAN card number',
                  },
               },
               {
                  className: 'col-6',
                  key: 'aadharcard',
                  type: 'input',
                  templateOptions: {
                     label: 'Aadhar card',
                     placeholder: 'Enter Aadhar card',
                  },
               },
            ],
         },
      ];
   }
}
