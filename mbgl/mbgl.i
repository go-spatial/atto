%module mbgl

%{
#include <mbgl/map/map.hpp>
#include <mbgl/util/image.hpp>
#include <mbgl/util/default_styles.hpp>
#include <mbgl/gl/headless_frontend.hpp>
#include <mbgl/util/default_thread_pool.hpp>
#include <mbgl/storage/default_file_source.hpp>
#include <mbgl/style/style.hpp>
%}

%include <mbgl/map/map.hpp>
%include <mbgl/util/image.hpp>
%include <mbgl/util/default_styles.hpp>
%include <mbgl/gl/headless_frontend.hpp>
%include <mbgl/util/default_thread_pool.hpp>
%include <mbgl/storage/default_file_source.hpp>
%include <mbgl/style/style.hpp>
