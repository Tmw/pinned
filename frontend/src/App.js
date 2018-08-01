import React, { Component } from "react";
import { Provider } from "mobx-react";
import Main from "./views/Main";
import Store from "./state";

class App extends Component {
  render() {
    return (
      <Provider store={Store}>
        <Main />
      </Provider>
    );
  }
}

export default App;
