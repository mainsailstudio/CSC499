import { Component, OnInit } from '@angular/core';
import { WordArray } from '../../api/api.constants';
import { UserConstantsService } from './../user-constants/user-constants.service';

@Component({
  selector: 'app-usability-test',
  templateUrl: './usability-test.component.html',
  styleUrls: ['./usability-test.component.css']
})
export class UsabilityTestComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-usability-test-instructions',
  templateUrl: './instructions.component.html',
  styleUrls: ['./usability-test.component.css']
})
export class UsabilityTestInstructionsComponent implements OnInit {

  // passphrase XKCD comic, great for illustrations
  private xkcdImage = '/assets/password_strength.png';

  // random word array for suggested words
  randomWordArray = [];
  lockArray = [];
  lockString = '';
  keyArray = [];

  // get the user's testLevel to know what to present them with
  testLevel = this.userConstants.TestLevel;

  // the default instruction ngSwitch state
  instructions = 1;

  constructor(private userConstants: UserConstantsService) { }

  ngOnInit() {
    this.refreshWords();
  }

  refreshWords() {
    // first clear array
    this.randomWordArray = [];
    const numArray = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    this.lockArray = [];
    this.lockString = '';
    this.keyArray = [];

    // then insert words
    // for this specific test, the user will always want around 10 words so this is fine
    for (let i = 0; i < 10; i++) {
      this.randomWordArray.push(WordArray[Math.floor(Math.random() * 2627)]);
    }

    // shuffle num array and take the first 4 for locks
    this.shuffle(numArray);
    this.lockArray = numArray.slice(0, 4);

    // put each random word associated with lock array into key array
    for (let i = 0; i < this.lockArray.length; i++) {
      this.keyArray.push(this.randomWordArray[this.lockArray[i] - 1]);
    }

    // make a nice string out of the lockArray to present
    this.lockString = this.lockArray.join(' - ');
  }

  changeInstructions(num: number) {
    if (num === 4) {
      console.log('Finished');
      return;
    }
    this.instructions = num;
  }

  /**
 * Shuffles array in place.
 * @param {Array} a items An array containing the items.
 */
  shuffle(a) {
    let j, x, i;
    for (i = a.length - 1; i > 0; i--) {
        j = Math.floor(Math.random() * (i + 1));
        x = a[i];
        a[i] = a[j];
        a[j] = x;
    }
  }


}
