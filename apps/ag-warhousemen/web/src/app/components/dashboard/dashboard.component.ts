import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { ToastrService } from 'ngx-toastr';
import { ConfigService } from 'src/app/services/config.service';

@Component({
   selector: 'app-dashboard',
   templateUrl: './dashboard.component.html',
   styleUrls: ['./dashboard.component.scss'],
})
export class DashboardComponent implements OnInit {
   identityData: any;

   constructor(
      private http: HttpClient,
      private spinner: NgxSpinnerService,
      private toastr: ToastrService,
      private configService: ConfigService
   ) {}

   ngOnInit(): void {
      this.fetch();
   }

   refresh() {
      this.fetch();
   }

   fetch() {
      this.spinner.show();

      this.http
         .get<any>(`${this.configService.baseUrl()}/api/v1/profile`)
         .subscribe({
            next: (v) => {
               const resp = JSON.parse(v['data']);
               this.identityData = JSON.stringify(resp, null, 2);
            },
            error: (e) => {
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => this.spinner.hide(),
         });
   }
}
