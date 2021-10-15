/* This example requires Tailwind CSS v2.0+ */
import { Disclosure } from '@headlessui/react'
import { SearchIcon, ShoppingBagIcon, ShoppingCartIcon, UserIcon } from '@heroicons/react/outline'

function classNames(...classes: any[]) {
    return classes.filter(Boolean).join(' ')
}

export default function UserNavbar() {
    return (
        <Disclosure as="nav" className="border-b border-gray-200">
            {({ open }) => (
                <>
                    <div className="container px-4 mx-auto sm:px-0">
                        <div className="flex items-center justify-between h-[60px]">
                            <div className="flex items-center space-x-10">
                                {/* Logo section */}
                                <div className="flex-shrink-0">
                                    <ShoppingBagIcon className="w-8 h-8 text-blue-700" />
                                </div>
                                <ul className="flex items-center space-x-10 font-semibold tracking-wide">
                                    <li>
                                        <a href="#">Testing</a>
                                    </li>
                                    <li>
                                        <a href="#">Testing</a>
                                    </li>
                                    <li>
                                        <a href="#">Testing</a>
                                    </li>
                                    <li>
                                        <a href="#">Testing</a>
                                    </li>
                                </ul>
                            </div>
                            <div className="flex items-center">
                                <div className="flex items-center space-x-5">
                                    <svg
                                        className="w-6 h-6"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth={1.5}
                                            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                                        />
                                    </svg>
                                    <svg
                                        className="w-6 h-6"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth={1.5}
                                            d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                                        />
                                    </svg>
                                </div>
                                <div className="w-px h-full mx-5 border border-gray-400">
                                    <span className="text-transparent">1</span>
                                </div>
                                <div className="flex items-center">
                                    <svg
                                        className="w-6 h-6"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth={1.5}
                                            d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                                        />
                                    </svg>
                                    <span className="text-lg font-semibold proportional-nums">0</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </>
            )}
        </Disclosure>
    )
}
