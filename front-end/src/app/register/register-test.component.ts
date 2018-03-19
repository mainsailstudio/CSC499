import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { trigger, state, transition, style, animate } from '@angular/animations';
import { Observable } from 'rxjs/Observable';

import { RegisterTestService } from './register-test.service';
import { LockKey } from './lock-key/lock-key';
import { Router } from '@angular/router';

export interface RegisterTestUser {
    id:         string;
    fname:      string;
    lname:      string;
    email:      string;
    init:       boolean;
    testLevel:  string;
    token:      string;
}

@Component({
  selector: 'app-register-test',
  templateUrl: 'register-test.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterTestComponent implements OnInit {

  register: RegisterTestUser[];

  constructor(private registerService: RegisterTestService, private router: Router) { }

  ngOnInit() { }

  testAPI() {
      const testUser = this.registerService.testAPI().subscribe(
          suc => {
            console.log(suc);
          },
          err => {
              console.log(err );
          }
      );
  }

  registerTestUser(email: string): void {
    this.register = undefined;
    email = email.trim();
    if (!email) { return; }

    // The server will generate the id for this new hero
    const testUser: RegisterTestUser = { email } as RegisterTestUser;
    this.registerService.registerTest(testUser).subscribe(
      suc => {
        Observable.timer(1000)
        .subscribe(i => {
          this.router.navigate(['/dashboard']);
        });
      },
      err => {
          console.log(err );
      }
    );
    // this.directUsers(gotUser);
  }

  // directUsers(testUser: RegisterTestUser): void {

  //   if (!testUser.init) {
  //       console.log('User not initalized');
  //       this.router.navigateByUrl('/dashboard');
  //   }
  // }

}
