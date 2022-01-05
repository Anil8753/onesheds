import { Component, OnInit } from '@angular/core';
import { ConfigService } from 'src/app/services/config.service';
import { ToastrService } from 'ngx-toastr';
import { HttpClient } from '@angular/common/http';
import { UserRegistrationData } from './types';

@Component({
   selector: 'app-profile',
   templateUrl: './profile.component.html',
   styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit {
   active = 1;
   regData!: UserRegistrationData;

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
               this.regData = JSON.parse(v['data']) as UserRegistrationData;
               console.log(this.regData);
            },
            error: (e) => {
               this.toastr.error('Failed to fetch the profile data.', 'Error!');
               console.error(e);
            },
            complete: () => console.info('complete'),
         });
   }
}
