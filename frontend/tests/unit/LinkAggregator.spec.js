import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import LinkAggregator from '@/components/LinkAggregator.vue'

describe('LinkAggregator.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(LinkAggregator, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders news', () => {
    const wrapper = shallowMount(LinkAggregator)
    wrapper.setData({ news_theme_list: [
      { 
        'title': 'Title', 
        'date': 'Date', 
        'url': 'URL' 
      }
    ] })
    expect(wrapper.text()).to.include('Title')
    expect(wrapper.text()).to.include('Date')
    expect(wrapper.text()).to.include('URL')
  })
})


