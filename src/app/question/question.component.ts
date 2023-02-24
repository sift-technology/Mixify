import { MixService } from './../mix.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css'],
})

export class QuestionComponent{
    title = "Questions: ";
    disabled = false;
    max = 100;
    min = 0;
    showTicks = false;
    step = 1;
    thumbLabel = true;
    value = 50;

    constructor(private myservice: MixService, private router: Router) {}

    public RESULTS: number[] = []

    onSubmit(data1: number, data2: number, data3: number, data4: number) {
      this.router.navigateByUrl('/results');
      console.log(data1, data2, data3, data4);
      this.RESULTS[0] = data1;
      this.RESULTS[1] = data2;
      this.RESULTS[2] = data3;
      this.RESULTS[3] = data4;
      this.myservice.getList(this.RESULTS);
      console.log(this.RESULTS);
    }
}
