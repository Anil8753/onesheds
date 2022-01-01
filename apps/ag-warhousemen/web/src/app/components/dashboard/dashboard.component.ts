import { Component, OnInit } from '@angular/core';
import { ConfigService } from 'src/app/services/config.service';
import { ToastrService } from 'ngx-toastr';
import { HttpClient } from '@angular/common/http';

@Component({
   selector: 'app-dashboard',
   templateUrl: './dashboard.component.html',
   styleUrls: ['./dashboard.component.scss'],
})
export class DashboardComponent implements OnInit {
   profileData = {};

   constructor(
      private http: HttpClient,
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
      this.http
         .get<any>(`${this.configService.baseUrl()}/api/v1/profile`)
         .subscribe({
            next: (v) => {
               const cert = JSON.parse(v['data']);
               this.profileData = JSON.stringify(cert, null, 4);
            },
            error: (e) => {
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => console.info('complete'),
         });
   }
}
