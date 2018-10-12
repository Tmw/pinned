import React from "react";
import { inject, observer } from "mobx-react";

import { Transition, animated } from "react-spring";

import ProgressIndicator from "components/ProgressIndicator";
import PinBox from "components/PinBox";
import OptionBox from "containers/OptionBox";
import PageTitle from "components/PageTitle";
import "stylesheets/pin-view.css";

class PinView extends React.Component {
  get challengeStore() {
    return this.props.store.ChallengeStore;
  }

  optionClickedHandler = optionID => {
    this.challengeStore.currentChallenge.answer(optionID);
  };

  render() {
    const challenge = this.challengeStore.currentChallenge,
      challengeIndex = this.challengeStore.currentChallengeIndex + 1,
      numChallenges = this.challengeStore.challenges.length;

    return (
      <React.Fragment>
        <PageTitle subtitle={`pin ${challengeIndex} / ${numChallenges}`} />
        <div className="view pin-view">
          <Transition
            native
            keys={challenge.id}
            from={{ opacity: 0, transform: "translateX(200px)" }}
            enter={{ opacity: 1, transform: "translateX(0px)" }}
            leave={{ opacity: 0, transform: "translateX(-200px)" }}
          >
            {style => (
              <animated.div
                className="pin-wrapper"
                key={challenge.id}
                style={{ ...style }}
              >
                <ProgressIndicator
                  step={challengeIndex}
                  total={numChallenges}
                />
                <PinBox text={challenge.text} />
                <OptionBox
                  options={challenge.choices}
                  onOptionClicked={this.optionClickedHandler}
                />
              </animated.div>
            )}
          </Transition>
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(PinView));
