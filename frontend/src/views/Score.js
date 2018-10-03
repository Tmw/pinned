import { inject, observer } from "mobx-react";
import React from "react";
import { Views } from "Constants";
import EmojiToUnicode from "components/EmojiToUnicode";
import WithSlackLinks from "components/WithSlackLinks";

import PageTitle from "components/PageTitle";

class ScoreView extends React.Component {
  get challengeStore() {
    return this.props.store.ChallengeStore;
  }

  get viewStore() {
    return this.props.store.ViewStore;
  }

  handleRetryClicked = () => {
    this.challengeStore.fetchChallenges();
    this.viewStore.presentView(Views.LOADING);
  };

  render() {
    const score = this.challengeStore.score;
    const numChallenges = this.challengeStore.challenges.length;

    return (
      <React.Fragment>
        <PageTitle subtitle="Your report card!" />

        <div className="score-view">
          <div className="score--box">
            <h1>
              <span role="img" aria-label="raised hands">
                🙌
              </span>
            </h1>
            <h1>
              {score} out of {numChallenges}!
            </h1>

            <ul>
              {this.challengeStore.challenges.map((c, i) => (
                <li
                  key={i}
                  className={`report-item ${c.isCorrect ? "green" : "red"}`}
                >
                  <WithSlackLinks>{EmojiToUnicode(c.text)}</WithSlackLinks>
                </li>
              ))}
            </ul>
            <br />
            <br />
            <br />

            <div className="zigzag" />

            <div className="center">
              <button
                className="button button--big"
                onClick={this.handleRetryClicked}
              >
                Try again
              </button>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(ScoreView));
