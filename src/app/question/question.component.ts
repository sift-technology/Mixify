import { MixService } from './../mix.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';

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

    constructor(private myservice: MixService) {}

    public RESULTS: number[] = []

    onSubmit(data1: string, data2: string, data3: string, data4: string) {
      console.log(data1, data2, data3, data4);
      this.RESULTS[0] = +data1;
      this.RESULTS[1] = +data2;
      this.RESULTS[2] = +data3;
      this.RESULTS[3] = +data4;

      this.myservice.getList(this.RESULTS)
    }
}
