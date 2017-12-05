// Must set up jsdom before importing react
import jsdom from 'jsdom';
const { JSDOM } = jsdom;
const { document } = (new JSDOM('<!doctype html><html><body></body></html>')).window;
global.document = document;
global.window = document.defaultView;

import React from 'react';
import { mount } from 'enzyme';
import sinon from 'sinon';
import sinonStubPromise from 'sinon-stub-promise';
sinonStubPromise(sinon);

import latencyFixtures from './fixtures/latencyTs.json';
import podFixtures from './fixtures/pods.json';
import Deployments from '../js/components/Deployments.jsx';

describe('Deployments', () => {
  let component, resp, fetchStub;

  function withPromise(fn) {
    return component.get(0).serverPromise.then(fn);
  }

  beforeEach(() => {
    fetchStub = sinon.stub(window, 'fetch');
  });

  afterEach(() => {
    window.fetch.restore();
  });

  it('renders the spinner before metrics are loaded', () => {
    fetchStub.returnsPromise().resolves({
      json: () => Promise.resolve({ metrics: [] })
    });
    component = mount(<Deployments />);

    expect(component.find("Deployments")).to.have.length(1);
    expect(component.find("ConduitSpinner")).to.have.length(1);
    expect(component.find("CallToAction")).to.have.length(0);
  });

  it('renders a call to action if no metrics are received', () => {
    fetchStub.returnsPromise().resolves({
      json: () => Promise.resolve({ metrics: [] })
    });
    component = mount(<Deployments />);

    return withPromise(() => {
      expect(component.find("Deployments")).to.have.length(1);
      expect(component.find("ConduitSpinner")).to.have.length(0);
      expect(component.find("CallToAction")).to.have.length(1);
    });
  });

  it('renders the deployments page if metrics are received', () => {
    fetchStub.returnsPromise().resolves({
      json: () => Promise.resolve({ metrics: [], pods: podFixtures.pods })
    });
    component = mount(<Deployments />);

    return withPromise(() => {
      expect(component.find("Deployments")).to.have.length(1);
      expect(component.find("ConduitSpinner")).to.have.length(0);
      expect(component.find("CallToAction")).to.have.length(0);
      expect(component.find("TabbedMetricsTable")).to.have.length(1);
    });
  });
});
