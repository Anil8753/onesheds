import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-question',
   templateUrl: './question.component.html',
   styleUrls: ['./question.component.scss'],
})
export class QuestionComponent implements OnInit {
   //
   @Input() item: any;
   warehouseId = '';
   knowledgebase: any[] = [];
   asked = '';

   constructor(
      public http: HttpClient,
      public configService: ConfigService,
      public activeModal: NgbActiveModal,
      private toastr: ToastrService
   ) {}

   ngOnInit(): void {
      this.warehouseId = this.item.warehouseId;
      this.fetch();
   }

   async fetch() {
      const url = `${this.configService.baseUrl()}/api/v1/knowledgebase/${
         this.item.warehouseId
      }`;

      const resp = await this.http.get<FAQResp>(url).toPromise();
      const data = JSON.parse(resp?.data as string) as IFAQ;
      this.knowledgebase = data.knowledgebase;
   }

   async ask() {
      if (this.asked === '') {
         this.toastr.error(`Please enter your question`);
         return;
      }

      const url = `${this.configService.baseUrl()}/api/v1/knowledgebase/question`;
      const postData = {
         warehouseId: this.item.warehouseId,
         question: this.asked,
      };

      this.http.post(url, postData).subscribe({
         next: (resp) => {
            this.asked = '';
            console.log(resp);
            this.toastr.success(`Your question placed suceesfully.`);
            this.fetch();
         },
         error: (e) => {
            this.toastr.error(`Question is not placed. error: ${e.message}`);
         },
      });
   }
}

interface FAQResp {
   data: string;
}

interface IFAQ {
   knowledgebase: any[];
}
