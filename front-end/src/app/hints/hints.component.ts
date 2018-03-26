import { Component, OnInit } from '@angular/core';
import { PracticeService } from '../dashboard/practice/practice.service';
import { UserConstantsService } from '../dashboard/user-constants/user-constants.service';
import { WordArray } from '../api/api.constants';

@Component({
  selector: 'app-dashboard-main-hints',
  templateUrl: './hints.component.html',
  styleUrls: ['./hints.component.css']
})
export class HintsComponent implements OnInit {

  keys = [];
  password = '';

  // image path stuff
  private relativePath = '../../assets/';
  private sharkImage = 'rainbow-shark.png';

  // random word array for suggested words
  randomWordArray = [];
  lockArray = [];
  lockString = '';
  keyArray = [];

  constructor(private practiceService: PracticeService,
              private userConstants: UserConstantsService) { }

  ngOnInit() {
    if (this.userConstants.TestLevel === 1) {
      this.practiceService.getTestUserDisplayPassword(this.userConstants.ID).subscribe(
        suc => {
          this.password = suc;
        },
        err => {
        }
      );
    }

    if (this.userConstants.TestLevel === 2 || this.userConstants.TestLevel === 3) {
      this.refreshWords();

      this.practiceService.getTestUserKeys(this.userConstants.ID).subscribe(
        suc => {
          this.keys = Array.from(suc); // what in the hell is going on with this, idk internet code works
        },
        err => {
        }
      );
    }

  }

  imagePath(image: string): string {
    const newImage = this.relativePath + image;
    return newImage;
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
