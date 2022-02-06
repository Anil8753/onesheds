import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-booking',
   templateUrl: './booking.component.html',
   styleUrls: ['./booking.component.scss'],
})
export class BookingComponent implements OnInit {
   @Input() item: any;
   constructor(
      private http: HttpClient,
      private toastr: ToastrService,
      private configService: ConfigService,
      public activeModal: NgbActiveModal
   ) {}

   ngOnInit(): void {}

   book() {
      const postData: NewOrderPostData = {
         WarehouseId: this.item.warehouseId,
      };

      this.http
         .post<any>(`${this.configService.baseUrl()}/api/v1/order`, postData)
         .subscribe({
            next: () => {
               this.toastr.success(`new order placed suceesfully.`);
            },
            error: (e) => {
               this.toastr.error(`new order is failed. error: ${e.message}`);
            },
         });
   }
}

interface NewOrderPostData {
   WarehouseId: string;
   // DepositorId: string;
   // FromDate    time.Time `json:"fromDate" binding:"required"`
}
