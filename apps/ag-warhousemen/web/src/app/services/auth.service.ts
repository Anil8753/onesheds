import { HttpClient, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { ConfigService } from './config.service';
import { INTERCEPTOR_NO_AUTH_HEADER } from '../services/constants.service';
import { Router } from '@angular/router';

const ACCESS_TOKEN_LOCAL_STORAGE_KEY = 'auth.accessToken';
// const REFRESH_TOKEN_LOCAL_STORAGE_KEY = 'auth.refreshToken';

@Injectable({
   providedIn: 'root',
})
export class AuthService {
   cachedRequests: Array<HttpRequest<any>> = [];

   constructor(
      private router: Router,
      private http: HttpClient,
      private configService: ConfigService
   ) {}

   public signIn(reqData: any): Observable<SigninResp> {
      return this.http
         .post<SigninResp>(
            `${this.configService.baseUrl()}/api/v1/signin`,
            reqData,
            {
               headers: { [INTERCEPTOR_NO_AUTH_HEADER]: 'true' },
            }
         )
         .pipe(
            tap((resp) => {
               this.setAccessToken(resp.data.accessToken);
            })
         );
   }

   logout(): void {
      this.removeAccessToken();
      this.router.navigateByUrl('/signin');
   }

   refresh(): Observable<SigninResp> {
      const postData = { refreshToken: this.getRefreshToken() };

      return this.http
         .post<SigninResp>(
            `${this.configService.baseUrl()}/api/v1/refreshToken`,
            postData,
            {
               headers: { [INTERCEPTOR_NO_AUTH_HEADER]: 'true' },
            }
         )
         .pipe(
            tap((resp) => {
               this.setAccessToken(resp.data.accessToken);
            })
         );
   }

   public removeAccessToken() {
      localStorage.removeItem(ACCESS_TOKEN_LOCAL_STORAGE_KEY);
   }

   public setAccessToken(token: string) {
      localStorage.setItem(ACCESS_TOKEN_LOCAL_STORAGE_KEY, token);
   }

   public getAccessToken(): string | null {
      return localStorage.getItem(ACCESS_TOKEN_LOCAL_STORAGE_KEY);
   }

   public getRefreshToken(): string | null {
      return '';
   }
}

export interface SigninResp {
   code: number;
   codeDesc: string;
   data: Data;
}

export interface Data {
   accessToken: string;
}
