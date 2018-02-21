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

  email = 'rando';

  public users;

  constructor(private _testService: TestService) { }

  ngOnInit() {
    this.getUsers();
  }

  getUsers() {
    this._testService.getUsers().subscribe(
      // the first argument is a function which runs on success
      data => { this.users = data; },
      // the second argument is a function which runs on error
      err => console.error(err),
      // the third argument is a function which runs on completion
      () => console.log('done loading users')
    );
  }

}
