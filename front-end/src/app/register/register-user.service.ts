import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { StartRegisterUser } from './register.component';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    // 'Authorization': 'my-auth-token' --------------- add here
  })
};

@Injectable()
export class RegisterUserService {
  registerUrl = 'http://localhost:8080/register';  // URL to web api
  private handleError: HandleError;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('RegisterStartService');
  }

//   /** GET heroes from the server */
//   getHeroes (): Observable<RegisterUser[]> {
//     return this.http.get<RegisterUser[]>(this.registerUrl)
//       .pipe(
//         catchError(this.handleError('getHeroes', []))
//       );
//   }

//   /* GET heroes whose name contains search term */
//   searchHeroes(term: string): Observable<Hero[]> {
//     term = term.trim();

//     // Add safe, URL encoded search parameter if there is a search term
//     const options = term ?
//      { params: new HttpParams().set('name', term) } : {};

//     return this.http.get<Hero[]>(this.heroesUrl, options)
//       .pipe(
//         catchError(this.handleError<Hero[]>('searchHeroes', []))
//       );
//   }

  //////// Save methods //////////

  /** POST: add a new user to the database */
  addUser (register: StartRegisterUser): Observable<StartRegisterUser> {
    return this.http.post<StartRegisterUser>(this.registerUrl, register, httpOptions)
      .pipe(
        catchError(this.handleError('registerUser', register))
      );
  }

//   /** DELETE: delete the hero from the server */
//   deleteHero (id: number): Observable<{}> {
//     const url = `${this.heroesUrl}/${id}`; // DELETE api/heroes/42
//     return this.http.delete(url, httpOptions)
//       .pipe(
//         catchError(this.handleError('deleteHero'))
//       );
//   }

//   /** PUT: update the hero on the server. Returns the updated hero upon success. */
//   updateHero (hero: Hero): Observable<Hero> {
//     httpOptions.headers =
//       httpOptions.headers.set('Authorization', 'my-new-auth-token');

//     return this.http.put<Hero>(this.heroesUrl, hero, httpOptions)
//       .pipe(
//         catchError(this.handleError('updateHero', hero))
//       );
//   }
}
