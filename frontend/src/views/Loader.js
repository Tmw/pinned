import React from "react";
import Spinner from "components/Spinner";

const getRandomLoadingPhrase = () => {
  const phrases = [
    "Hold on a moment..",
    "Preparing...",
    "Un momento, por favor!",
    "Crunching the numbers..",
    "Fetching pins",
    "Doing Maths"
  ];

  const randomIndex = Math.floor(Math.random() * phrases.length);
  return phrases[randomIndex];
};

export default class LoaderView extends React.Component {
  render() {
    return (
      <div className="view loader-view">
        <Spinner text={getRandomLoadingPhrase()} />
      </div>
    );
  }
}
