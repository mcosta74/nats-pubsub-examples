import { Injectable, OnDestroy } from '@angular/core';

import { Subject, Observable } from 'rxjs';

import { connect, NatsConnection, StringCodec } from 'nats.ws';

export interface ChatMessage {
  subject: string;
  data?: string;
}

@Injectable({
  providedIn: 'root',
})
export class NatsService implements OnDestroy {
  private conn: NatsConnection | null;

  private _messages$ = new Subject<ChatMessage>();
  get messages$(): Observable<ChatMessage> {
    return this._messages$.asObservable();
  }

  get isConnected(): boolean {
    return this.conn !== null;
  }

  constructor() {
    this.conn = null;
  }

  async ngOnDestroy() {
    if (this.conn !== null) {
      await this.conn.close();
    }
  }

  async connect() {
    if (this.isConnected) {
      return;
    }

    try {
      this.conn = await connect({
        servers: 'wss://massimo-mbp.fwx.one:8443',
        // servers: 'demo.nats.io',

        name: 'Web App',
      });
    } catch (err) {
      console.error('Error:', err);
      return;
    }

    const sc = StringCodec();
    const sub = this.conn.subscribe("chat.*");

    for await (const m of sub) {
      this._messages$.next({subject: m.subject, data: sc.decode(m.data)});
    }
  }
}
