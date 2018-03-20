import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';
import { APIURL, TestURL } from '../api/api.constants';
import { LoginActivity, ConfigActivity } from './log.interface';


@Injectable()
export class ActivityLogService {

  private handleError: HandleError;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('ActivityLogService');
  }

  // log a user succesfully configuring their account in
  logConfigActivity(log: ConfigActivity) {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    return this.http.post<ConfigActivity>(APIURL + 'test/log-config', log, httpOptionsAuthorized)
      .pipe(
        catchError(this.handleError('configActivityLog', log))
      );
  }

  // log a user succesfully logging in
  // NOTE: the testLevel variable tells if it is a password or dynauth
  logLoginActivity(log: LoginActivity) {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    return this.http.post<LoginActivity>(APIURL + 'test/log-login', log, httpOptionsAuthorized)
      .pipe(
        catchError(this.handleError('loginActivityLog', log))
      );
  }

}
