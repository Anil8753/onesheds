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
import { getDirtyValues } from 'src/app/utils/form';

@Component({
   selector: 'app-wh-infra',
   templateUrl: './wh-infra.component.html',
   styleUrls: ['./wh-infra.component.scss'],
})
export class WhInfraComponent implements OnChanges, OnInit {
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
      const changes = getDirtyValues(this.form);

      this.http
         .put<any>(`${this.configService.baseUrl()}/api/v1/warehouse`, {
            warehouseId: this.data.warehouseId,
            properties: changes,
         })
         .subscribe({
            next: (v) => {
               this.form.markAsPristine();
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
                  key: 'yearOfConstruction',
                  type: 'number',
                  templateOptions: {
                     label: 'Year Of Construction (YYYY)',
                     placeholder: 'Year of construction (1984)',
                     min: 1700,
                     max: 2030,
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'constructionStandard',
                  type: 'select',
                  defaultValue: 'NA',
                  templateOptions: {
                     label: 'Construction Standard',
                     //  multiple: true,
                     required: true,
                     options: [
                        { label: 'NA', value: 'NA' },
                        { label: 'BIS', value: 'BIS' },
                        { label: 'CWC', value: 'CWC' },
                        { label: 'FCI', value: 'FCI' },
                        { label: 'NHB', value: 'NHB' },
                        { label: 'NHM', value: 'NHM' },
                        { label: 'SHM', value: 'SHM' },
                        { label: 'NCCD', value: 'NCCD' },
                        { label: 'MoFPI', value: 'MoFPI' },
                        { label: 'APEDA', value: 'APEDA' },
                        {
                           label: 'State Govt. Norms',
                           value: 'State Govt. Norms',
                        },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'constructionNormColdStorage',
                  type: 'input',
                  templateOptions: {
                     placeholder: 'Cold storage construction norm',
                     label: 'Construction Norm for Cold Storage',
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'entryExitGatesGaurds',
                  type: 'select',
                  defaultValue: 'No',
                  templateOptions: {
                     label: 'Entry/Exit Gates Manned by Guards',
                     required: true,
                     options: [
                        { label: 'Yes', value: 'Yes' },
                        { label: 'No', value: 'No' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'noOfSecurityGaurds',
                  type: 'number',
                  templateOptions: {
                     label: 'No of Security Gaurds',
                     placeholder: 'Enter Security Gaurds',
                     required: true,
                     min: 0,
                     max: 100,
                  },
               },
               {
                  className: 'col-4',
                  key: 'fireHydrants',
                  type: 'select',
                  defaultValue: 'No',
                  templateOptions: {
                     label: 'Fire Hydrant, Static Water Tank',
                     required: true,
                     options: [
                        { label: 'Yes', value: 'Yes' },
                        { label: 'No', value: 'No' },
                     ],
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'fireSaftyAlarm',
                  type: 'select',
                  defaultValue: 'No',
                  templateOptions: {
                     label: 'Fire Safety Alarms available',
                     required: true,
                     options: [
                        { label: 'Yes', value: 'Yes' },
                        { label: 'No', value: 'No' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'nightLight',
                  type: 'select',
                  defaultValue: 'No',
                  templateOptions: {
                     label: 'Adequate Night Light Arrangment',
                     required: true,
                     options: [
                        { label: 'Yes', value: 'Yes' },
                        { label: 'No', value: 'No' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'noOfFireBuckets',
                  type: 'number',
                  templateOptions: {
                     label: 'Number of Fire Buckets',
                     placeholder: 'Fire Buckets',
                     required: true,
                     min: 0,
                     max: 10000,
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-4',
                  key: 'lorryWeighBridge',
                  type: 'select',
                  defaultValue: 'Outside',
                  templateOptions: {
                     label: 'Lorry Weighbridge',
                     required: true,
                     options: [
                        { label: 'Inside', value: 'Inside' },
                        { label: 'Outside', value: 'Outside' },
                     ],
                  },
               },
               {
                  className: 'col-4',
                  key: 'lorryWeighBridgeDistance',
                  type: 'input',
                  templateOptions: {
                     label: 'Lorry Weighbridge Distance (KM)',
                     placeholder: 'Lorry Weighbridge Distance (KM)',
                  },
               },
               {
                  className: 'col-4',
                  key: 'lorryWeighBridgeArea',
                  type: 'textarea',
                  templateOptions: {
                     rows: 2,
                     label: 'Lorry Weighbridge Address',
                     placeholder: 'Lorry Weighbridge Address',
                  },
               },
            ],
         },
      ];
   }
}

export const DummyModel = () => {
   return {};
};
