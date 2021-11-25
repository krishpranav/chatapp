
import './App.css';
import {Component} from "react";
import { connect, sendMsg} from "./api/apirequest";

class ChatApplication extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("HEllo")
    sendMsg("Helllo");
  }

  render() {
    return (
        <div className="App">
          <button onClick={this.send}>HIT</button>
        </div>
    );
  }
}

export default ChatApplication;