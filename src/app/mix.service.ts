import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MixService {

  constructor(private http: HttpClient) { }

  getList(results:number[]) {
    
  }
}
