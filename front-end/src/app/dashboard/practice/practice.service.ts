import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../../http-error-handler.service';
import { APIURL, TestURL } from '../../api/api.constants';


@Injectable()
export class PracticeService {

  private handleError: HandleError;
  keys = [];

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('PracticeService');
  }

  getTestUserKeys(userid: string): Observable<string[]> {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        // authorize with token here
      })
    };

    const params = new HttpParams().set('userID', userid); // create new HttpParams

    return this.http.get<string[]>(APIURL + 'test/get-keys', { headers: httpOptionsAuthorized.headers, params: params });
  }

  getTestUserDisplayPassword(userid: string): Observable<string> {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        // authorize with token here
      })
    };

    const params = new HttpParams().set('userID', userid); // create new HttpParams

    return this.http.get<string>(APIURL + 'test/get-pass', { headers: httpOptionsAuthorized.headers, params: params });
  }

}
