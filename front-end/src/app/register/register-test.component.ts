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
  showSuccess = false;
  showFail = false;
  showInit = false;

  constructor(private registerService: RegisterTestService, private router: Router) { }

  ngOnInit() { }

  registerTestUser(email: string): void {
    this.showSuccess = false;
    this.showFail = false;

    this.register = undefined;
    email = email.trim();
    if (!email) { return; }

    const testUser: RegisterTestUser = { email } as RegisterTestUser;
    this.registerService.registerTest(testUser).subscribe(
      suc => {
        if (suc == null) {
          this.showSuccess = false;
          this.showFail = false;
          this.showInit = true;
          return;
        }
        this.showSuccess = true;
        this.showFail = false;
        this.showInit = false;

        Observable.timer(1000)
        .subscribe(i => {
          this.router.navigate(['/dashboard']);
        });
      },
      err => {
        this.showSuccess = false;
        this.showFail = true;
        this.showInit = false;
      }
    );
  }

}
