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
   selector: 'app-wh-surrounding',
   templateUrl: './wh-surrounding.component.html',
   styleUrls: ['./wh-surrounding.component.scss'],
})
export class WhSurroundingComponent implements OnChanges, OnInit {
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
                  className: 'col-6',
                  key: 'jurisdictionPoliceStation',
                  type: 'input',
                  templateOptions: {
                     label: 'Name of jurisdiction Police Station',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'jurisdictionPoliceStationDistance',
                  type: 'number',
                  templateOptions: {
                     label: 'Police Station Distance',
                     placeholder: 'Enter Distance (KM)',
                     min: 0,
                     max: 200,
                  },
               },
            ],
         },
         {
            fieldGroupClassName: 'row',
            fieldGroup: [
               {
                  className: 'col-6',
                  key: 'fireStation',
                  type: 'input',
                  templateOptions: {
                     label: 'Name of Fire Station',
                     required: true,
                  },
               },
               {
                  className: 'col-4',
                  key: 'fireStationDistance',
                  type: 'number',
                  templateOptions: {
                     label: 'Fire Station Distance',
                     placeholder: 'Enter Distance (KM)',
                     min: 0,
                     max: 200,
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
