import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-review',
   templateUrl: './review.component.html',
   styleUrls: ['./review.component.scss'],
})
export class ReviewComponent implements OnInit {
   //
   @Input() item: any;
   warehouseId = '';
   reviews: any[] = [];

   constructor(
      public http: HttpClient,
      public configService: ConfigService,
      public activeModal: NgbActiveModal
   ) {}

   ngOnInit(): void {
      this.warehouseId = this.item.warehouseId;
      this.fetchReview();
   }

   async fetchReview() {
      const url = `${this.configService.baseUrl()}/api/v1/review/warehouse/${
         this.item.warehouseId
      }`;

      const resp = await this.http.get<ReviewResp>(url).toPromise();
      const data = JSON.parse(resp?.data as string);
      this.reviews = data;
   }

   async postReview() {
      const url = `${this.configService.baseUrl()}/api/v1/review`;

      const postData = {
         warehouseId: this.warehouseId,
         userRating: 4.5,
         reviewText: 'This is test review',
      };

      await this.http.post(url, postData).toPromise();
      await this.fetchReview();
   }
}

interface ReviewResp {
   data: string;
}

// interface IReview {
//    faqs: any[];
// }
