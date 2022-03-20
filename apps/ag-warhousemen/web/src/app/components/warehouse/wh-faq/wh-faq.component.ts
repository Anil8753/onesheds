import { HttpClient } from '@angular/common/http';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';

import {
   Component,
   Input,
   OnChanges,
   OnInit,
   SimpleChanges,
} from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-wh-faq',
   templateUrl: './wh-faq.component.html',
   styleUrls: ['./wh-faq.component.scss'],
})
export class WhFaqComponent implements OnInit, OnChanges {
   @Input() data: any;
   name: string = '';
   whId: string = '';
   faqs: any[] = [];

   faq = {
      question: '',
      answer: '',
   };

   constructor(
      private http: HttpClient,
      private spinner: NgxSpinnerService,
      private toastr: ToastrService,
      private configService: ConfigService,
      private modalService: NgbModal
   ) {}

   ngOnInit(): void {
      this.init(this.data);
   }

   ngOnChanges(changes: SimpleChanges): void {
      for (const propName in changes) {
         const change = changes[propName];
         this.init(change.currentValue);
      }
   }

   init(currData: any) {
      this.whId = currData.warehouseId;

      this.name = currData.properties.warehousename
         ? (currData.properties.warehousename as string)
         : 'New Warehouse';

      this.fetchFAQs();
   }

   openDlg(content: any) {
      this.modalService.open(content, { ariaLabelledBy: 'modal-basic-title' });
   }

   onAdd() {
      this.addFAQ();
   }

   onDelete(index: number) {
      this.deleteFAQ(index);
   }

   fetchFAQs() {
      this.http
         .get<any>(
            `${this.configService.baseUrl()}/api/v1/faq/warehouse/${this.whId}`
         )
         .subscribe({
            next: (v) => {
               const items = JSON.parse(v.data);
               const faqs = items.faqs;
               if (!!faqs && faqs.length > 0) {
                  this.faqs = faqs;
               }
               this.spinner.hide();
            },
            error: (e) => {
               this.spinner.hide();
               this.toastr.error('Failed to fetch the faq data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }

   addFAQ() {
      if (this.faq.question === '' || this.faq.answer === '') {
         return;
      }

      interface AddFAQDataPostData {
         warehouseId: string;
         question: string;
         input: string;
      }

      const postData: AddFAQDataPostData = {
         warehouseId: this.whId,
         question: this.faq.question,
         input: `{ "text":"${this.faq.answer}", "extra":"https://google.com" }`,
      };

      this.spinner.show();

      this.http
         .post<any>(`${this.configService.baseUrl()}/api/v1/faq/add`, postData)
         .subscribe({
            next: (v) => {
               this.faq.answer = '';
               this.faq.question = '';

               this.spinner.hide();
               this.modalService.dismissAll();
               this.fetchFAQs();
            },
            error: (e) => {
               this.spinner.hide();
               this.toastr.error('Failed to add faq data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }

   deleteFAQ(i: number) {
      interface DeleteFAQDataPostData {
         warehouseId: string;
         index: number;
      }

      const postData: DeleteFAQDataPostData = {
         warehouseId: this.whId,
         index: i,
      };

      this.spinner.show();

      this.http
         .put<any>(
            `${this.configService.baseUrl()}/api/v1/faq/delete`,
            postData
         )
         .subscribe({
            next: (v) => {
               this.spinner.hide();
               this.fetchFAQs();
            },
            error: (e) => {
               this.spinner.hide();
               this.toastr.error('Failed to delete faq data.', 'Error!');
               console.error(e);
            },
            complete: () => {},
         });
   }
}
