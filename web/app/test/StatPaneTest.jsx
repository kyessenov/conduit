import React from 'react';
import { mount } from 'enzyme';
import StatPane from '../js/components/StatPane.jsx';

describe('StatPane', () => {
  it('renders the request, success rate and latency components', () => {
    let component = mount(<StatPane
      lastUpdated={Date.now()}
      summaryMetrics={[]} />);

    expect(component.find(".border-container").length).to.equal(2);
    expect(component.find(".line-graph").length).to.equal(3);
  });
});
