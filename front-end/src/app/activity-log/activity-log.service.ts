import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';
import { APIURL, TestURL } from '../api/api.constants';
import { LogActivity } from './log.interface';


@Injectable()
export class ActivityLogService {

  private handleError: HandleError;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('ActivityLogService');
  }

  // log a user succesfully logging in
  // NOTE: the testLevel variable tells if it is a password or dynauth
  logActivity(log: LogActivity) {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    return this.http.post<LogActivity>(APIURL + 'test/log-activity', log, httpOptionsAuthorized)
      .pipe(
        catchError(this.handleError('logActivity', log))
      );
  }

}
