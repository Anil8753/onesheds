import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-faq',
   templateUrl: './faq.component.html',
   styleUrls: ['./faq.component.scss'],
})
export class FaqComponent implements OnInit {
   //
   @Input() item: any;
   warehouseId = '';
   faqs: any[] = [];

   constructor(
      public http: HttpClient,
      public configService: ConfigService,
      public activeModal: NgbActiveModal
   ) {}

   ngOnInit(): void {
      this.warehouseId = this.item.warehouseId;
      this.fetchFAQ();
   }

   async fetchFAQ() {
      const url = `${this.configService.baseUrl()}/api/v1/faq/${
         this.item.warehouseId
      }`;

      const resp = await this.http.get<FAQResp>(url).toPromise();
      const data = JSON.parse(resp?.data as string) as IFAQ;
      this.faqs = data.faqs;
   }
}

interface FAQResp {
   data: string;
}

interface IFAQ {
   faqs: any[];
}
