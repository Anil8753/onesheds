import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';
import { BookingComponent } from '../booking/booking.component';

@Component({
   selector: 'app-listing',
   templateUrl: './listing.component.html',
   styleUrls: ['./listing.component.scss'],
})
export class ListingComponent implements OnInit {
   ps = 2;
   items: any[] = [];
   bookmark: string = '';

   constructor(
      private http: HttpClient,
      private toastr: ToastrService,
      private configService: ConfigService,
      private modalService: NgbModal
   ) {}

   ngOnInit(): void {
      this.fetchWarehousePage(this.bookmark);
   }

   fetchWarehousePage(bm: string) {
      const q = `{"selector":{"docType": "WarehouseRegData"}}`;

      this.http
         .get<any>(
            `${this.configService.baseUrl()}/api/v1/warehouse/querypagination?q=${q}&pagesize=${
               this.ps
            }&bookmark=${bm}`
         )
         .subscribe({
            next: (v) => {
               const data: PaginationQueryResult = JSON.parse(v['data']);
               this.bookmark = data.bookmark;
               this.items = data.records;
            },
            error: (e) => {
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => console.info('complete'),
         });
   }

   next() {
      this.fetchWarehousePage(this.bookmark);
   }

   format(r: any) {
      return JSON.stringify(r, null, 4);
   }

   book(item: any) {
      const modalRef = this.modalService.open(BookingComponent, {
         size: 'lg',
         backdrop: 'static',
      });
      modalRef.componentInstance.item = item;
   }
}

export interface PaginationQueryResult {
   records: any[];
   fetchedRecordsCount: number;
   bookmark: string;
}
