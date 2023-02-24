import { Injectable } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MixService {

  constructor(private http: HttpClient) { }

  getList(results:number[]) {
    return this.http.post<any>('http://localhost:8080/results', results);
  }
}
