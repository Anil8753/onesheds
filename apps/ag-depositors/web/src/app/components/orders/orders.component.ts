import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-orders',
   templateUrl: './orders.component.html',
   styleUrls: ['./orders.component.scss'],
})
export class OrdersComponent implements OnInit {
   items: any[] = [];
   constructor(
      private http: HttpClient,
      private toastr: ToastrService,
      private configService: ConfigService
   ) {}

   ngOnInit(): void {}

   fetchOrders() {
      this.http
         .get<any>(`${this.configService.baseUrl()}/api/v1/orders`)
         .subscribe({
            next: (v) => {
               this.items = v;
            },
            error: (e) => {
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => console.info('complete'),
         });
   }
}
