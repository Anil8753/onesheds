import { Injectable } from '@angular/core';

export const INTERCEPTOR_NO_AUTH_HEADER = 'INTERCEPTOR_NO_AUTH_HEADER';
const ACCESS_TOKEN_LOCAL_STORAGE_KEY = 'auth.accessToken';
const REFRESH_TOKEN_LOCAL_STORAGE_KEY = 'auth.refreshToken';

@Injectable({
   providedIn: 'root',
})
export class ConstantsService {
   constructor() {}
}
