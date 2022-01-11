import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

import { ToastrService } from 'ngx-toastr';
import { NgxSpinnerService } from 'ngx-spinner';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-warehouse',
   templateUrl: './warehouse.component.html',
   styleUrls: ['./warehouse.component.scss'],
})
export class WarehouseComponent implements OnInit {
   active = '1';

   warehouses: any[] = [];
   curWarehouse: any;
   curDisplayed: string = '';

   constructor(
      private http: HttpClient,
      private spinner: NgxSpinnerService,
      private toastr: ToastrService,
      private configService: ConfigService
   ) {}

   ngOnInit(): void {
      this.fetch();
   }

   onChange(i: number) {
      this.setCurItem(this.warehouses[i]);
   }

   onNew() {
      this.spinner.show();

      this.http
         .post<any>(`${this.configService.baseUrl()}/api/v1/warehouse`, {})
         .subscribe({
            next: (v) => {
               console.log(v);
               this.toastr.success('Created successfully.', 'Success!');

               const item = JSON.parse(v.data);
               this.warehouses.push(item);

               this.setCurItem(item);
            },
            error: (e) => {
               this.toastr.error('Failed to create.', 'Error!');
               console.error(e);
            },
            complete: () => this.spinner.hide(),
         });
   }

   onUpdate(data: any) {
      console.log(data);
      this.setCurItem(data);
   }

   displayName(data: any): string {
      if (!data || !data.properties || !data.properties.warehousename) {
         return 'New Warehouse';
      }

      return data.properties.warehousename;
   }

   setCurItem(item: any) {
      this.curWarehouse = item;
      this.curDisplayed = this.displayName(this.curWarehouse);
   }

   private fetch() {
      this.spinner.show();

      this.http
         .get<any>(`${this.configService.baseUrl()}/api/v1/warehouse`)
         .subscribe({
            next: (v) => {
               const items = JSON.parse(v.data);
               if (!!items && items.length > 0) {
                  this.warehouses = items;
                  this.setCurItem(items[0]);
               }
               this.spinner.hide();
            },
            error: (e) => {
               this.spinner.hide();
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }
}
