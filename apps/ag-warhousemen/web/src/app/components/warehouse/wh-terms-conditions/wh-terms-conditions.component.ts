import { HttpClient } from '@angular/common/http';
import { NgxSpinnerService } from 'ngx-spinner';
import AWN from 'awesome-notifications';

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
import { FormArray, FormBuilder, FormControl, FormGroup } from '@angular/forms';

@Component({
   selector: 'app-wh-terms-conditions',
   templateUrl: './wh-terms-conditions.component.html',
   styleUrls: ['./wh-terms-conditions.component.scss'],
})
export class WhTermsConditionsComponent implements OnChanges, OnInit {
   @Output() updateEvent = new EventEmitter<any>();

   @Input() data: any;

   form = this.fb.group({
      clauses: this.fb.array([this.fb.control('')]),
   });

   constructor(
      private fb: FormBuilder,
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

   get list() {
      return this.form.get('clauses') as FormArray;
   }

   onAdd() {
      this.list.push(this.fb.control(''));
   }

   onRemove(i: number) {
      const ctrl = this.list.controls[i];
      if (ctrl.value.length === 0) {
         this.list.removeAt(i);
         return;
      }

      new AWN().confirm(
         'Are you sure you want to delete?',
         () => {
            this.list.removeAt(i);
            this.form.markAsDirty();
         },
         () => {},
         {
            labels: { confirm: 'Delete' },
         }
      );
   }
   private init(data: any) {
      this.data = data;

      if (!!this.data && !!this.data.termsConditions) {
         this.list.clear();
         const list = this.data.termsConditions as string[];
         list.forEach((l) => this.list.push(this.fb.control(l)));
      }
   }

   submit() {
      const clauses = this.list.controls
         .map((ctrl) => ctrl.value)
         .filter((v) => !!v && v.length);

      this.spinner.show();

      this.http
         .put<any>(`${this.configService.baseUrl()}/api/v1/warehouse`, {
            warehouseId: this.data.warehouseId,
            termsConditions: clauses,
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
}
