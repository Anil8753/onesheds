import { Component } from '@angular/core';
import { FieldArrayType } from '@ngx-formly/core';
import AWN from 'awesome-notifications';

@Component({
   selector: 'formly-repeat-section',
   template: ` <div
         *ngFor="let field of field.fieldGroup; let i = index"
         class="row"
      >
         <formly-field class="col" [field]="field"></formly-field>
         <div class="col-sm-2 d-flex align-items-center">
            <button class="btn btn-danger" type="button" (click)="onRemove(i)">
               Remove
            </button>
         </div>
      </div>
      <div style="margin:10px 0;">
         <button class="btn btn-primary" type="button" (click)="add()">
            {{ to['addText'] }}
         </button>
      </div>`,
})
export class RepeatTypeComponent extends FieldArrayType {
   onRemove(i: number) {
      let notifier = new AWN();

      let onOk = () => {
         this.remove(i);
      };

      let onCancel = () => {};

      notifier.confirm('Are you sure?', onOk, onCancel, {
         labels: {
            confirm: 'Dangerous action',
         },
      });
   }
}
