import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  constructor() { }

  public baseUrl() : string {
    return 'http://localhost:9000';
  }
}
