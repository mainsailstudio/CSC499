import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../../http-error-handler.service';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    // 'Authorization': 'my-auth-token' --------------- add here
  })
};

@Injectable()
export class ConfigurationService {
  registerUrl = 'http://13.92.156.114:8080/register';  // URL to web api
  private handleError: HandleError;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('RegisterStartService');
  }

//    /** PUT: update user in database */
//    continueUser (register: ContinueRegisterUser): Observable<ContinueRegisterUser> {
//     return this.http.post<ContinueRegisterUser>(this.registerUrl, register, httpOptions)
//       .pipe(
//         catchError(this.handleError('registerUser', register))
//       );
//   }

//   /** PUT: update user in database */
//   finalUser (register: FinalRegisterUser): Observable<FinalRegisterUser> {
//     return this.http.post<FinalRegisterUser>(this.registerUrl, register, httpOptions)
//       .pipe(
//         catchError(this.handleError('registerUser', register))
//       );
//   }
}

