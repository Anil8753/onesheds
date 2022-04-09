import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit, SimpleChanges } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-wh-knowledgebase',
   templateUrl: './wh-knowledgebase.component.html',
   styleUrls: ['./wh-knowledgebase.component.scss'],
})
export class WhKnowledgebaseComponent implements OnInit {
   @Input() data: any;

   whId: string = '';
   knowledgebase: any[] = [];

   answerText = '';

   constructor(
      private http: HttpClient,
      private spinner: NgxSpinnerService,
      private toastr: ToastrService,
      private configService: ConfigService
   ) {}

   ngOnInit(): void {}

   ngOnChanges(changes: SimpleChanges): void {
      for (const propName in changes) {
         const change = changes[propName];
         this.init(change.currentValue);
      }
   }

   init(currData: any) {
      this.whId = currData.warehouseId;
      this.fetchKnowledgeBase();
   }

   fetchKnowledgeBase() {
      //
      this.knowledgebase = [];

      this.http
         .get<any>(
            `${this.configService.baseUrl()}/api/v1/knowledgebase/${this.whId}`
         )
         .subscribe({
            next: (v) => {
               const items = JSON.parse(v.data);
               const knowledgebase = items.knowledgebase;
               if (!!knowledgebase && knowledgebase.length > 0) {
                  this.knowledgebase = knowledgebase;
               }
               this.spinner.hide();
            },
            error: (e) => {
               this.spinner.hide();
               //  this.toastr.error('Failed to fetch the knowlwdgebase data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }

   onAddAnswer(i: number) {
      if (this.answerText === '') {
         return;
      }

      interface PostData {
         warehouseId: string;
         index: number;
         input: string;
      }

      const postData: PostData = {
         warehouseId: this.whId,
         index: i,
         input: `{ "text":"${this.answerText}", "extra":"https://yahoo.com" }`,
      };

      this.spinner.show();

      this.http
         .post<any>(
            `${this.configService.baseUrl()}/api/v1/knowledgebase/answer`,
            postData
         )
         .subscribe({
            next: (v) => {
               this.answerText = '';

               this.spinner.hide();
               this.fetchKnowledgeBase();
            },
            error: (e) => {
               this.spinner.hide();
               this.toastr.error('Failed to add knowledgebase data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }
}
