import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Observable} from 'rxjs/Observable';
import {TestService} from './test.service';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent implements OnInit {

  displayRegisterStart = true;
  displayRegisterContinue = false;
  displayRegisterFinal = false;
  public users;

 // constructor(private _testService: TestService) { }
  constructor() { }

  ngOnInit() {
    // this.getUsers();
  }

  register(): void {
    if (this.displayRegisterStart === true) {
      this.displayRegisterStart = false;
      this.displayRegisterContinue = true;
    } else if (this.displayRegisterContinue === true) {
        this.displayRegisterContinue = false;
        this.displayRegisterFinal = true;

    }
  }

  // getUsers() {
  //   this._testService.getUsers().subscribe(
  //     // the first argument is a function which runs on success
  //     data => { this.users = data; },
  //     // the second argument is a function which runs on error
  //     err => console.error(err),
  //     // the third argument is a function which runs on completion
  //     () => console.log('done loading users')
  //   );
  // }

}
