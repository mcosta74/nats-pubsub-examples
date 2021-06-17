import { Component, OnDestroy, OnInit } from '@angular/core';

import { Subscription } from 'rxjs';
import '@cds/core/button/register';

import { NatsService, ChatMessage } from '../nats.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit, OnDestroy {

  messages: ChatMessage[] = [];

  private _subs: Subscription;

  constructor(private nats: NatsService) {
    this._subs = Subscription.EMPTY;
  }

  ngOnInit(): void {
    this._subs = this.nats.messages$.subscribe(
      (msg) => (this.messages = [msg, ...this.messages.slice(0, 20)])
    );
  }

  ngOnDestroy(): void {
    this._subs.unsubscribe();
  }

  onConnect() {
    this.nats.connect();
  }

  get isConnected(): boolean {
    return this.nats.isConnected;
  }
}
