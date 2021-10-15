import React from 'react'
import { act } from 'react-dom/test-utils'
import { mount, ReactWrapper, shallow } from 'enzyme'
import InputForm from './InputForm'

describe('<InputForm />', () => {
  const smiles = 'CCC(CC)CC'
  const errorMessages = {
    required: 'SMILES is required',
    wrongFormat: "SMILES must consists of only 'C' and round brackets",
    startWithC: "SMILES must start with 'C'",
  }

  it('should submit given valid data', async () => {
    const handleSubmit = jest.fn()
    const wrapper: ReactWrapper = mount(
      <InputForm smiles={smiles} onSubmit={handleSubmit} />
    )
    const input: ReactWrapper = wrapper.find('input[name="smiles"]')
    const button: ReactWrapper = wrapper.find('button[type="submit"]')

    expect(input).toHaveLength(1)
    expect(button).toHaveLength(1)

    expect(input.at(0).prop('value')).toEqual(smiles)

    await act(async () => {
      wrapper.find('form').simulate('submit')
    })

    expect(handleSubmit).toHaveBeenCalledWith(smiles)

    wrapper.unmount()
  })

  it('should render validation error given empty smiles', async () => {
    const handleSubmit = jest.fn()
    const wrapper: ReactWrapper = mount(
      <InputForm smiles="" onSubmit={handleSubmit} />
    )
    const input: ReactWrapper = wrapper.find('input[name="smiles"]')

    await act(async () => {
      input.simulate('change', {
        target: {
          name: 'smiles',
          value: '',
        },
      })
      input.simulate('blur')
    })
    await act(async () => {
      wrapper.find('form').simulate('submit')
    })

    expect(handleSubmit).not.toHaveBeenCalled()
    expect(wrapper.find('div.invalid-feedback').text()).toEqual(
      errorMessages['required']
    )
  })

  it('should render validation error given smiles of wrong format', async () => {
    const handleSubmit = jest.fn()
    const wrapper: ReactWrapper = mount(
      <InputForm smiles="" onSubmit={handleSubmit} />
    )
    const input: ReactWrapper = wrapper.find('input[name="smiles"]')

    await act(async () => {
      input.simulate('change', {
        target: {
          name: 'smiles',
          value: 'DD(CC)EF',
        },
      })
      input.simulate('blur')
    })

    await act(async () => {
      wrapper.find('form').simulate('submit')
    })
    expect(wrapper.find('div.invalid-feedback').text()).toEqual(
      errorMessages['wrongFormat']
    )
  })

  it("should render validation error given smiles which doesn't start with C", async () => {
    const handleSubmit = jest.fn()
    const wrapper: ReactWrapper = mount(
      <InputForm smiles="" onSubmit={handleSubmit} />
    )
    const input: ReactWrapper = wrapper.find('input[name="smiles"]')

    await act(async () => {
      input.simulate('change', {
        target: {
          name: 'smiles',
          value: '(C)CCC',
        },
      })
      input.simulate('blur')
    })

    await act(async () => {
      wrapper.find('form').simulate('submit')
    })

    expect(wrapper.find('div.invalid-feedback').text()).toEqual(
      errorMessages['startWithC']
    )
  })
})
