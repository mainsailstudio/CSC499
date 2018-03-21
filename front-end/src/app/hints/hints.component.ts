import { Component, OnInit } from '@angular/core';
import { PracticeService } from '../dashboard/practice/practice.service';
import { UserConstantsService } from '../dashboard/user-constants/user-constants.service';

@Component({
  selector: 'app-dashboard-main-hints',
  templateUrl: './hints.component.html',
  styleUrls: ['./hints.component.css']
})
export class HintsComponent implements OnInit {

  keys = [];

  constructor(private practiceService: PracticeService,
              private userConstants: UserConstantsService) { }

  ngOnInit() {
    this.practiceService.getTestUserKeys(this.userConstants.ID).subscribe(
      suc => {
        console.log(suc);
        this.keys = Array.from(suc); // what in the hell is going on with this, idk internet code works
        console.log('Base suc is ' + suc);
        console.log('This keys is ' + this.keys);
      },
      err => {
        console.log(err);
        console.log('Error log here');
      }
    );
  }
}
