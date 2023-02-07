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

    onSubmit(data: string, data2: string, data3: string, data4: string) {
      console.log(data, data2, data3, data4);
    }
}
