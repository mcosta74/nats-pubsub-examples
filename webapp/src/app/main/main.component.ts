import { Component, OnDestroy, OnInit } from '@angular/core';

import '@cds/core/button/register';
import { NatsService } from '../nats.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {

  constructor(private nats: NatsService) {
    this.nats.messages$.subscribe((msg) => console.log(msg));
  }

  ngOnInit(): void {
  }


  onConnect() {
    console.log("Connect");
    this.nats.connect();
  }

}
