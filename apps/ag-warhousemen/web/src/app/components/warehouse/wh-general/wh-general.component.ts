import { HttpClient } from '@angular/common/http';
import { NgxSpinnerService } from 'ngx-spinner';

import {
   Component,
   EventEmitter,
   Input,
   OnChanges,
   OnInit,
   Output,
   SimpleChanges,
} from '@angular/core';

import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';
import { FormGroup } from '@angular/forms';
import { FormlyFormOptions, FormlyFieldConfig } from '@ngx-formly/core';

@Component({
   selector: 'app-wh-general',
   templateUrl: './wh-general.component.html',
   styleUrls: ['./wh-general.component.scss'],
})
export class WhGeneralComponent implements OnChanges, OnInit {
   @Output() updateEvent = new EventEmitter<any>();

   @Input() data: any;
   model = {};
   form = new FormGroup({});
   options: FormlyFormOptions = {};
   fields: FormlyFieldConfig[] = [];

   constructor(
      private http: HttpClient,
      private toastr: ToastrService,
      private spinner: NgxSpinnerService,
      private configService: ConfigService
   ) {}

   ngOnInit(): void {}

   ngOnChanges(changes: SimpleChanges) {
      for (const propName in changes) {
         const change = changes[propName];
         console.log(change);
         this.init(change.currentValue);
      }
   }

   private init(data: any) {
      this.data = data;

      if (!!this.data && !!this.data.properties) {
         this.model = JSON.parse(JSON.stringify(this.data.properties));
      } else {
         this.model = DummyModel();
      }

      this.initFormly();
   }

   submit() {
      if (!this.form.valid) {
         return;
      }

      this.spinner.show();

      this.http
         .put<any>(`${this.configService.baseUrl()}/api/v1/warehouse`, {
            warehouseId: this.data.warehouseId,
            properties: this.model,
         })
         .subscribe({
            next: (v) => {
               this.toastr.success('Updated successfully.', 'Success!');
               this.updateEvent.emit(JSON.parse(v.data));
            },
            error: (e) => {
               this.toastr.error('Failed to update.', 'Error!');
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
                  className: 'col-4',
                  key: 'ownershipType',
                  type: 'select',
                  templateOptions: {
                     required: true,
                     label: 'Ownership type',
                     options: [
                        { label: 'Owned', value: 'owned' },
                        { label: 'Leased', value: 'leased' },
                        { label: 'Hired', value: 'hired' },
                        { label: 'Rented', value: 'rented' },
                        { label: 'Sub Leased', value: 'sub_leased' },
                        { label: 'Revenue Sharing', value: 'revenue_sharing' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'warehouseType',
                  type: 'select',
                  templateOptions: {
                     required: true,
                     label: 'Warehouse type',
                     options: [
                        {
                           label: 'Conventional Warehouse',
                           value: 'conventional_warehouse',
                        },
                        {
                           label: 'Cold Storage Warehouse',
                           value: 'cold_storage_warehouse',
                        },
                        { label: 'Silos', value: 'silos' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'capacity',
                  type: 'input',
                  templateOptions: {
                     required: true,
                     label: 'Capacity (in MT)',
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'warehousename',
                  type: 'textarea',
                  templateOptions: {
                     label: 'Warehouse name',
                     placeholder: 'Enter Warehouse name',
                     rows: 1,
                     required: true,
                  },
               },
               {
                  className: 'col-8',
                  key: 'warehouseaddress',
                  type: 'textarea',
                  templateOptions: {
                     label: 'Address',
                     placeholder: 'Enter Warehouse Address',
                     rows: 1,
                     required: true,
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'district',
                  type: 'input',
                  templateOptions: {
                     label: 'District',
                     placeholder: 'Enter District',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'state',
                  type: 'input',
                  templateOptions: {
                     label: 'State',
                     placeholder: 'Enter State',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'pincode',
                  type: 'input',
                  templateOptions: {
                     label: 'PIN Code',
                     placeholder: 'Enter PIN Code',
                     required: true,
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'email',
                  type: 'input',
                  templateOptions: {
                     label: 'Email Address',
                     placeholder: 'Enter Email Address',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'contact',
                  type: 'input',
                  templateOptions: {
                     label: 'Contact number',
                     placeholder: 'Enter Contact number',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'contact2',
                  type: 'input',
                  templateOptions: {
                     label: 'Second Contact number',
                     placeholder: 'Enter Contact number',
                  },
               },
            ],
         },
      ];
   }
}

export const DummyModel = () => {
   return {
      ownershipType: '',
      warehouseType: '',
      capacity: '',
      warehousename: '',
      warehouseaddress: '',
      address: '',
      district: '',
      state: '',
      pincode: '',
      contact: '',
      contact2: '',
   };
};
