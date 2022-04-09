import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit, SimpleChanges } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-wh-review',
   templateUrl: './wh-review.component.html',
   styleUrls: ['./wh-review.component.scss'],
})
export class WhReviewComponent implements OnInit {
   //
   @Input() data: any;
   whId = '';
   reviews: any[] = [];

   constructor(
      public http: HttpClient,
      private toastr: ToastrService,
      public configService: ConfigService
   ) {}

   ngOnInit(): void {
      console.log(this.data);
   }

   ngOnChanges(changes: SimpleChanges): void {
      console.log(changes);
      for (const propName in changes) {
         const change = changes[propName];
         this.init(change.currentValue);
      }
   }

   init(currData: any) {
      this.whId = currData.warehouseId;
      this.fetchReviews();
   }

   async fetchReviews() {
      const url = `${this.configService.baseUrl()}/api/v1/review/${this.whId}`;

      const resp = await this.http.get<ReviewResp>(url).toPromise();
      const data = JSON.parse(resp?.data as string);
      this.reviews = data;
   }

   async postReview() {
      try {
         const url = `${this.configService.baseUrl()}/api/v1/review`;

         const testReview = this.getTestReview();
         const postData = {
            warehouseId: this.whId,
            userRating: testReview.userRating,
            reviewText: testReview.reviewText,
         };

         await this.http.post(url, postData).toPromise();
         await this.fetchReviews();
         this.toastr.success('review added successfully.', 'Success!');
      } catch (e) {
         this.toastr.error('Failed to add review.', 'Error!');
      }
   }

   async onAddReply(reviewId: string, targetId: string) {
      try {
         const url = `${this.configService.baseUrl()}/api/v1/review_reply`;

         const postData = {
            reviewId: reviewId,
            targetId: targetId,
            replyText: `this is test reply (${Math.floor(
               Math.random() * 100 + 1
            )})`,
         };

         await this.http.post(url, postData).toPromise();
         await this.fetchReviews();
         this.toastr.success('review reply added successfully.', 'Success!');
      } catch (e) {
         this.toastr.error('Failed to add review reply.', 'Error!');
      }
   }

   getTestReview() {
      const reviews = [
         { userRating: 4.5, reviewText: 'This is very good' },
         { userRating: 1.0, reviewText: 'This is very bad' },
         { userRating: 2.0, reviewText: 'front road is is narrow' },
         { userRating: 3.0, reviewText: 'This is ok, need improvements' },
      ];

      const i = Math.floor(Math.random() * 4 + 1);
      return reviews[i];
   }
}

interface ReviewResp {
   data: string;
}
