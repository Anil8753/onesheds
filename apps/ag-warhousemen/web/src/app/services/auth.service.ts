import { HttpClient, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { ConfigService } from './config.service';
import { INTERCEPTOR_NO_AUTH_HEADER } from '../services/constants.service';

const ACCESS_TOKEN_LOCAL_STORAGE_KEY = 'auth.accessToken';
// const REFRESH_TOKEN_LOCAL_STORAGE_KEY = 'auth.refreshToken';

@Injectable({
   providedIn: 'root',
})
export class AuthService {
   cachedRequests: Array<HttpRequest<any>> = [];

   constructor(
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

   public setAccessToken(token: string) {
      localStorage.setItem(ACCESS_TOKEN_LOCAL_STORAGE_KEY, token);
   }

   public getAccessToken(): string | null {
      return localStorage.getItem(ACCESS_TOKEN_LOCAL_STORAGE_KEY);
   }

   // public isAuthenticated(): boolean {
   //   // get the token
   //   const token = this.getToken() as string;
   //   // return a boolean reflecting
   //   // whether or not the token is expired
   //   return tokenNotExpired(token);
   // }

   public collectFailedRequest(request: HttpRequest<any>): void {
      this.cachedRequests.push(request);
   }
   public retryFailedRequests(): void {
      // retry the requests. this method can
      // be called after the token is refreshed
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
